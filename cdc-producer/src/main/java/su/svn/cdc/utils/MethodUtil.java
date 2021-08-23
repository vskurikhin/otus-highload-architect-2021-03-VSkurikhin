package su.svn.cdc.utils;

import com.google.common.base.CaseFormat;
import lombok.extern.slf4j.Slf4j;

import java.lang.reflect.Field;
import java.lang.reflect.Method;
import java.lang.reflect.Modifier;

@Slf4j
public class MethodUtil {
    public static final String GETTER = "get";
    public static final String SETTER = "set";
    public static final String IS = "is";

    public static boolean isValidGetter(Method method, Field field) {
        return method != null && method.getReturnType().equals(field.getType());
    }

    public static boolean isGetter(Method method) {
        return (method.getName().startsWith(GETTER) || method.getName().startsWith(IS))
                && method.getParameterCount() == 0
                && ! method.getReturnType().equals(void.class)
                && ! Modifier.isVolatile(method.getModifiers());
    }

    public static boolean isValidSetter(Method method, Field field) {
        return method != null && methodTypeCompatibleFieldType(method, field);
    }

    public static boolean methodTypeCompatibleFieldType(Method method, Field field) {
        Class<?> parameterType = method.getParameterTypes()[0];
        return parameterType.equals(field.getType())
                || (parameterType.equals(Boolean.class) && boolean.class.equals(field.getType()))
                || (parameterType.equals(Character.class) && char.class.equals(field.getType()))
                || (parameterType.equals(Byte.class) && byte.class.equals(field.getType()))
                || (parameterType.equals(Short.class) && short.class.equals(field.getType()))
                || (parameterType.equals(Integer.class) && int.class.equals(field.getType()))
                || (parameterType.equals(Long.class) && long.class.equals(field.getType()))
                || (parameterType.equals(Float.class) && float.class.equals(field.getType()))
                || (parameterType.equals(Double.class) && double.class.equals(field.getType()))
                || (parameterType.equals(boolean.class) && Boolean.class.equals(field.getType()))
                || (parameterType.equals(char.class) && Character.class.equals(field.getType()))
                || (parameterType.equals(byte.class) && Byte.class.equals(field.getType()))
                || (parameterType.equals(short.class) && Short.class.equals(field.getType()))
                || (parameterType.equals(int.class) && Integer.class.equals(field.getType()))
                || (parameterType.equals(long.class) && Long.class.equals(field.getType()))
                || (parameterType.equals(float.class) && Float.class.equals(field.getType()))
                || (parameterType.equals(double.class) && Double.class.equals(field.getType()));
    }

    public static boolean isSetter(Method method) {
        return method.getName().startsWith(SETTER)
                && method.getParameterCount() == 1
                && method.getReturnType().equals(void.class)
                && ! Modifier.isVolatile(method.getModifiers());
    }

    public static String name(String fieldName, String prefix) {
        int index = prefix.length();
        StringBuilder sb = new StringBuilder(prefix);
        sb.append(fieldName);
        sb.setCharAt(index, Character.toUpperCase(sb.charAt(index)));

        return sb.toString();
    }

    public static String convertToLowerCamelCase(String name) {
        return CaseFormat.LOWER_UNDERSCORE.to(CaseFormat.LOWER_CAMEL, name);
    }

    public static String convertToUpperCamelCase(String name) {
        return CaseFormat.LOWER_UNDERSCORE.to(CaseFormat.UPPER_CAMEL, name);
    }
}
