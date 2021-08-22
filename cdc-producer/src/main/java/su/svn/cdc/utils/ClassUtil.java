package su.svn.cdc.utils;

import lombok.extern.slf4j.Slf4j;
import org.apache.avro.specific.SpecificRecordBase;

import javax.annotation.Nullable;

@Slf4j
public class ClassUtil {

    @Nullable
    public static Class<? extends SpecificRecordBase> classForName(String name) {
        try {
            //noinspection unchecked
            return (Class<? extends SpecificRecordBase>) Class.forName(name);
        } catch (ClassNotFoundException e) {
            log.error("classForName ", e);
        }
        return null;
    }
}
