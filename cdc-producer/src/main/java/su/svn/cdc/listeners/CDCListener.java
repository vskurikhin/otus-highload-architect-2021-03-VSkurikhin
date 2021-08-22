package su.svn.cdc.listeners;

import lombok.extern.slf4j.Slf4j;
import org.apache.avro.specific.SpecificRecordBase;
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

import javax.annotation.PostConstruct;
import javax.annotation.PreDestroy;
import java.nio.ByteBuffer;
import java.util.Map;
import java.util.concurrent.Executor;
import java.util.concurrent.Executors;

import static io.debezium.connector.AbstractSourceInfo.DATABASE_NAME_KEY;
import static io.debezium.connector.AbstractSourceInfo.SCHEMA_NAME_KEY;
import static io.debezium.connector.AbstractSourceInfo.TABLE_NAME_KEY;
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

    // private final Map<String, String> destination;

//    private final Map<String, Class<? extends SpecificRecordBase>> mapToClasses;

//    private final BinaryMessageEncoder<Envelope> envelopeEncoder;
//
//    private final BinaryMessageEncoder<MetaEnvelope> metaEnvelopeEncoder;

    private final KafkaTemplate<String, ByteBuffer> kafkaTemplate;

    private final String topic;

    /**
     * Конструктор, который загружает конфигурацию и устанавливает метод обратного вызова handleEvent,
     * который вызывается при выполнении транзакционной операции базы данных.
     * @param abdbConnector
     * @param kafkaTemplate
     * @param topic
     */
    private CDCListener(
            Configuration abdbConnector,
            KafkaTemplate<String, ByteBuffer> kafkaTemplate,
            @Value("${app.cap-topic}") String topic) {
        this.engine = EmbeddedEngine
                .create()
                .using(abdbConnector)
                .notifying(this::handleEvent).build();
        this.kafkaTemplate = kafkaTemplate;
        this.topic = topic;
    }

    /**
     * The method is called after the Debezium engine is initialized and started asynchronously using the Executor.
     */
    @PostConstruct
    private void start() {
        this.executor.execute(engine);
    }

    /**
     * This method is called when the container is being destroyed. This stops the debezium, merging the Executor.
     */
    @PreDestroy
    private void stop() {
        if (this.engine != null) {
            this.engine.stop();
        }
    }

    /**
     * This method is invoked when a transactional action is performed on any of the tables that were configured.
     *
     * @param sourceRecord
     */
    private void handleEvent(SourceRecord sourceRecord) {
        Struct sourceRecordValue = (Struct) sourceRecord.value();
/*
        if (sourceRecordValue != null) {
            Operation operation = Operation.forCode((String) sourceRecordValue.get(OPERATION));

            if (operation != null) {
                String record = AFTER; // For Update & Insert operations.
                // Only if this is a transactional operation.
                if (operation != Operation.READ) {
                    if (operation == Operation.DELETE)
                        record = BEFORE; // For Delete operations.
                }
                sendMessage(sourceRecordValue, operation, record);
            } else*/
                log.info("sourceRecordValue: {}", sourceRecordValue);
        // }
    }

    private void sendMessage(Struct sourceRecordValue, Operation operation, String record) {

        final Struct sourceStruct = (Struct) sourceRecordValue.get(SOURCE);
        String db = sourceStruct.get(DATABASE_NAME_KEY).toString();
        String schema = sourceStruct.get(SCHEMA_NAME_KEY).toString();
        String table = sourceStruct.get(TABLE_NAME_KEY).toString();
        Struct afterStruct = (Struct) sourceRecordValue.get(record);
        Map<String, Object> data = afterStruct.schema().fields().stream()
                .map(Field::name)
                .filter(fieldName -> afterStruct.get(fieldName) != null)
                .map(fieldName -> Pair.of(fieldName, afterStruct.get(fieldName)))
                .collect(toMap(Pair::getKey, Pair::getValue));
        data.put("_operation", operation.code());
        log.trace("Data Changed: on database: {}, schema: {} on table: {} with Operation: {}. Message: {}",
                 db, schema, table, operation.name(), data);
//        try {
//            sendMessage(schema, table, operation, data);
//        } catch (IOException | ErrorCase exception) {
//            log.error("sendMessage ", exception);
//        }
    }
/*
    private void sendMessage(String schema, String table, Operation op, Map<String, Object> data) throws IOException {

        String key = schema + '.' + table;
        Class<? extends SpecificRecordBase> avroClass = mapToClasses.get(key);
        log.trace("key:{} -> {}", key, avroClass);

        if (avroClass != null) {
            ByteBuffer byteBuffer = getByteBuffer(avroClass, data);
            ByteBuffer crc32 = Crc32Util.crc32(byteBuffer);
            Envelope envelope = Envelope.newBuilder()
                    .setVersion(1)
                    .setSignAlg("CRC32")
                    .setSign(crc32)
                    .setMContainer(byteBuffer)
                    .setTypeName(destination.get(key))
                    .build();
            kafkaTemplate.send(topic, envelopeEncoder.encode(envelope));
            log.info("sent Envelope with key:{} and Operation:{}", key, op);
        } else
            log.error("for key:{} avroClass is null", key);
    }

    private <T extends SpecificRecordBase> ByteBuffer getByteBuffer(Class<T> tClass, Map<String, Object> message)
            throws IOException {
        log.trace("getByteBuffer({}, {})", tClass, message);
        return encoderFabric.get(tClass).encode(datumFabric.get(tClass, message));
    }
*/
}
