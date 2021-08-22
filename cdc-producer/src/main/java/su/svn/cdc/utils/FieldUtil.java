package su.svn.cdc.utils;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import su.svn.cdc.model.ReadOnlyConvert;

import java.lang.reflect.Field;
import java.lang.reflect.Modifier;
import java.math.BigDecimal;
import java.math.BigInteger;
import java.time.Duration;
import java.time.Instant;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;
import java.time.OffsetDateTime;
import java.time.OffsetTime;
import java.time.ZonedDateTime;
import java.util.Calendar;
import java.util.Currency;
import java.util.Date;
import java.util.TimeZone;
import java.util.UUID;
import java.util.function.Consumer;

public class FieldUtil {
    private static final Logger LOGGER = LoggerFactory.getLogger(FieldUtil.class);

    public static boolean isTransientOrStatic(Field field) {
        return Modifier.isTransient(field.getModifiers())
                || Modifier.isStatic(field.getModifiers());
    }

    public static boolean isSimpleObject(Object value) {
        return value instanceof Boolean
                || value instanceof Byte
                || value instanceof Character
                || value instanceof Short
                || value instanceof Integer
                || value instanceof Long
                || value instanceof Float
                || value instanceof Double
                || value instanceof BigInteger
                || value instanceof BigDecimal
                || value instanceof UUID
                || value instanceof String
                || value instanceof Class
                || value instanceof TimeZone
                || value instanceof Currency
                || value instanceof Date
                || value instanceof Calendar
                || value instanceof Duration
                || value instanceof Instant
                || value instanceof LocalDateTime
                || value instanceof LocalDate
                || value instanceof LocalTime
                || value instanceof OffsetDateTime
                || value instanceof OffsetTime
                || value instanceof ZonedDateTime;
        // TODO || value instanceof java.sql.Date
        // TODO || value instanceof Time
        // TODO || value instanceof Timestamp
    }

    public static  <T> void updateIfNotNull(Consumer<T> consumer, T o) {
        if (o != null) consumer.accept(o);
    }

    public static boolean isNotReadOnly(Field field) {
        boolean readOnly = field.isAnnotationPresent(ReadOnlyConvert.class);
        return ! readOnly;
    }
}
