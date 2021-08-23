package su.svn.cdc.config;

import lombok.extern.slf4j.Slf4j;
import org.springframework.boot.autoconfigure.kafka.KafkaProperties;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.kafka.core.DefaultKafkaProducerFactory;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.kafka.core.ProducerFactory;

import java.util.AbstractMap;
import java.util.Map;
import java.util.stream.Collectors;

@Slf4j
@Configuration
public class KafkaConfiguration {

    private final KafkaProperties kafkaProperties;

    public KafkaConfiguration(KafkaProperties kafkaProperties) {
        this.kafkaProperties = kafkaProperties;
    }

    @Bean
    public ProducerFactory<String, String> producerAuditFactory() {

        Map<String, Object> producerProperties = kafkaProperties.buildProducerProperties().entrySet().stream()
                .peek(e -> log.debug("KafkaProperties({}): {}", e.getKey(), e.getValue()))
                .map(e -> new AbstractMap.SimpleEntry<>(e.getKey(), e.getValue()))
                .collect(Collectors.toMap(Map.Entry::getKey, Map.Entry::getValue));
        log.debug("producerProperties: {}", kafkaProperties.getProperties().toString());

        return new DefaultKafkaProducerFactory<>(producerProperties);
    }

    @Bean
    public KafkaTemplate<String, String> kafkaTemplate() {
        return new KafkaTemplate<>(producerAuditFactory());
    }
}
