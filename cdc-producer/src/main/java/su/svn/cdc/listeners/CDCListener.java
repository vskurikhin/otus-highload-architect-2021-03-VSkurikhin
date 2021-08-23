package su.svn.cdc.listeners;

import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;

import io.debezium.config.Configuration;
import io.debezium.embedded.EmbeddedEngine;
import org.apache.commons.lang3.tuple.Pair;
import org.apache.kafka.connect.data.Field;
import org.apache.kafka.connect.data.Struct;
import org.apache.kafka.connect.source.SourceRecord;
import su.svn.cdc.utils.Operation;

import javax.annotation.Nonnull;
import javax.annotation.PostConstruct;
import javax.annotation.PreDestroy;
import java.io.IOException;
import java.util.AbstractMap;
import java.util.Map;
import java.util.concurrent.Executor;
import java.util.concurrent.Executors;
import java.util.stream.Collectors;

import static io.debezium.connector.AbstractSourceInfo.DATABASE_NAME_KEY;
import static io.debezium.connector.AbstractSourceInfo.TABLE_NAME_KEY;
import static io.debezium.data.Envelope.FieldName.AFTER;
import static io.debezium.data.Envelope.FieldName.BEFORE;
import static io.debezium.data.Envelope.FieldName.SOURCE;
import static java.util.stream.Collectors.toMap;

@Slf4j
@Component
public class CDCListener {

    public static final int CRC32_CAPACITY = 8;
    /**
     * Однопоточный пул, который будет запускать движок Debezium асинхронно.
     */
    private final Executor executor = Executors.newSingleThreadExecutor();

    /**
     * Движок Debezium, в который необходимо загрузить конфигурацию, и запустить - для работы CDC.
     */
    private final EmbeddedEngine engine;

    private final KafkaTemplate<String, String> kafkaTemplate;

    private final String topic;

    /**
     * Конструктор, который загружает конфигурацию и устанавливает метод обратного вызова handleEvent,
     * который вызывается при выполнении транзакционной операции базы данных.
     *
     * @param abdbConnector
     * @param kafkaTemplate
     * @param topic
     */
    private CDCListener(Configuration abdbConnector,
            KafkaTemplate<String, String> kafkaTemplate,
            @Value("${app.topic}") String topic) {
        this.engine = EmbeddedEngine
                .create()
                .using(abdbConnector)
                .notifying(this::handleEvent).build();
        this.kafkaTemplate = kafkaTemplate;
        this.topic = topic;
    }

    /**
     * Метод вызывается после инициализации движка Debezium и асинхронного запуска с использованием Executor.
     */
    @PostConstruct
    private void start() {
        this.executor.execute(engine);
    }

    /**
     * Этот метод вызывается при уничтожении контейнера. Это останавливает Debezium, сливая Executor.
     */
    @PreDestroy
    private void stop() {
        if (this.engine != null) {
            this.engine.stop();
        }
    }

    /**
     * Этот метод вызывается, когда транзакционное действие выполняется над любой из настроенных таблиц.
     *
     * @param sourceRecord - подготовленная структура (org.apache.kafka.connect.source) для передачи
     *                       в Kafka Connect.
     */
    private void handleEvent(SourceRecord sourceRecord) {

        Struct sourceRecordValue = (Struct) sourceRecord.value();

        if (sourceRecordValue != null) {
            log.trace("sourceRecordValue: {}", sourceRecordValue);
            var field = sourceRecordValue.schema().field("op");
            if (field != null) {
                var operation = operationForCode(sourceRecordValue.get(field));
                log.trace("operation: {}", operation);
                if (operation != null) {
                    checkOperationBeforeSend(sourceRecordValue, operation);
                } else
                    log.warn("operation is null");
            } else
                log.warn("field is null");
        } else
            log.warn("sourceRecordValue is null");
    }

    private Operation operationForCode(Object code) {
        return Operation.forCode(code.toString());
    }

    private void checkOperationBeforeSend(Struct sourceRecordValue, @Nonnull Operation operation) {

        String record = AFTER;  // Для операций Update & Insert operations.
                                // Только если это транзакционная операция.
        if (operation != Operation.READ) {
            if (operation == Operation.DELETE)
                record = BEFORE; // Для операции Delete.
        }
        final Struct sourceStruct = (Struct) sourceRecordValue.get(SOURCE);
        final Struct dataStruct = (Struct) sourceRecordValue.get(record);
        sendMessage(sourceStruct, dataStruct, operation);
    }

    private void sendMessage(Struct sourceStruct, Struct dataStruct, Operation operation) {

        String db = sourceStruct.get(DATABASE_NAME_KEY).toString();
        // String schema = sourceStruct.get(SCHEMA_NAME_KEY).toString();
        String table = sourceStruct.get(TABLE_NAME_KEY).toString();
        Map<String, Object> data = dataStruct.schema().fields().stream()
                .map(Field::name)
                .filter(fieldName -> dataStruct.get(fieldName) != null)
                .map(fieldName -> Pair.of(fieldName, dataStruct.get(fieldName)))
                .collect(toMap(Pair::getKey, Pair::getValue));
        data.put("_operation", operation.code());
        log.trace("Data Changed: on database: {}, on table: {} with Operation: {}. Message: {}",
                db, table, operation.name(), data);
        data = prepareMap(data);
        try {
            sendMessage(db, table, operation, data);
        } catch (IOException exception) {
            log.error("sendMessage ", exception);
        }
    }

    private Map<String, Object> prepareMap(Map<String, Object> data) {
        Object o = data.get("id");
        if (o instanceof Long) {
            long id = (long) o;
            data.put("id", Long.toUnsignedString(id));
        }
        return data.entrySet().stream()
                .peek(e -> log.trace("data({}): {}", e.getKey(), e.getValue()))
                .map(e -> new AbstractMap.SimpleEntry<>(up(e.getKey()), e.getValue()))
                .collect(Collectors.toMap(Map.Entry::getKey, Map.Entry::getValue));

    }

    private String up(String key) {
        return Character.toUpperCase(key.charAt(0)) + key.substring(1);
    }

    private void sendMessage(String db, String table, Operation op, Map<String, Object> data) throws IOException {

        String key = db + '.' + table;
        ObjectMapper mapper = new ObjectMapper();
        String json = mapper.writerWithDefaultPrettyPrinter().writeValueAsString(data);
        kafkaTemplate.send(topic, key, json);
        log.info("sent Message with key:{} and Operation:{} and json:{}", key, op, json);
    }
}
