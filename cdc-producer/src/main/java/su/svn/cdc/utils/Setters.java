package su.svn.cdc.utils;

import lombok.extern.slf4j.Slf4j;

import java.lang.reflect.Field;
import java.lang.reflect.Method;
import java.util.Arrays;
import java.util.Collections;
import java.util.ConcurrentModificationException;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
import java.util.function.BiConsumer;
import java.util.function.Function;
import java.util.stream.Collectors;

@Slf4j
public class Setters {

    public static final boolean FIELD_SNAKE_CASE = false;

    private static final String META_SETTERS_SUFFIX = "$1";

    private final Map<String, Method> methodMap;

    private final Map<String, BiConsumer<Object, Object>> biConsumerMap;

    public Setters(Class<?> aClass) {
        this.methodMap = Collections.unmodifiableMap(setters(aClass, FIELD_SNAKE_CASE));
        this.biConsumerMap = Collections.unmodifiableMap(setterBiConsumers(methodMap));
    }

    public Setters(Class<?> aClass, boolean fieldSnakeCase) {
        this.methodMap = Collections.unmodifiableMap(setters(aClass, fieldSnakeCase));
        this.biConsumerMap = Collections.unmodifiableMap(setterBiConsumers(methodMap));
    }

    public Method get(String name) {
        return methodMap.get(name);
    }

    public BiConsumer<Object, Object> getBiConsumer(String name) {
        return biConsumerMap.get(name);
    }


    public void forEach(BiConsumer<String, BiConsumer<Object, Object>> action) {
        Objects.requireNonNull(action);
        for (Map.Entry<String, BiConsumer<Object, Object>> entry : biConsumerMap.entrySet()) {
            String k;
            BiConsumer<Object, Object> v;
            try {
                k = entry.getKey();
                v = entry.getValue();
            } catch(IllegalStateException ise) {
                // this usually means the entry is no longer in the map.
                throw new ConcurrentModificationException(ise);
            }
            action.accept(k, v);
        }
    }

    private static Map<String, Method> setters(Class<?> aClass, boolean fieldSnakeCase) {
        Map<String, Method> methods = Arrays.stream(aClass.getMethods())
                .filter(su.svn.cdc.utils.MethodUtil::isSetter)
                .collect(Collectors.toMap(Method::getName, Function.identity()));
        Map<String, Method> result = new HashMap<>();

        for (Field field : aClass.getDeclaredFields()) {
            if ( ! FieldUtil.isTransientOrStatic(field) && FieldUtil.isNotReadOnly(field)) {
                String fieldName = fieldSnakeCase ? su.svn.cdc.utils.MethodUtil.convertToUpperCamelCase(field.getName()) : field.getName();
                String methodName = su.svn.cdc.utils.MethodUtil.name(fieldName, su.svn.cdc.utils.MethodUtil.SETTER);
                if (field.getName().charAt(0) == '_') {
                    //this field is for meta information in avro schemas and has specific setter
                    methodName = methodName + META_SETTERS_SUFFIX;
                }
                Method setter = methods.get(methodName);
                if (MethodUtil.isValidSetter(setter, field)) {
                    result.put(field.getName(), setter);
                }
            }
        }
        return result;
    }

    private static Map<String, BiConsumer<Object, Object>> setterBiConsumers(Map<String, Method> map) {
        Map<String, BiConsumer<Object, Object>> result = new HashMap<>();
        for(Map.Entry<String, Method> entry : map.entrySet()) {
            result.put(entry.getKey(), setterBiConsumer(entry.getValue()));
        }
        return result;
    }

    private static BiConsumer<Object, Object> setterBiConsumer(Method setter) {
        return (o, value) -> {
            try {
                Class<?> valueClass = value != null ? value.getClass() : null;
                log.trace("setterBiConsumer: {}.{}({}:{})", o.getClass(), setter.getName(), valueClass, value);
                setter.invoke(o, value);
            } catch (Exception e) {
                logError("Setter invoke exception: ", e);
            }
        };
    }

    private static void logError(String message, Throwable throwable) {
        log.error(message, throwable);
    }

    private static void logTrace(Object o) {
        log.trace(o.toString());
    }
}
