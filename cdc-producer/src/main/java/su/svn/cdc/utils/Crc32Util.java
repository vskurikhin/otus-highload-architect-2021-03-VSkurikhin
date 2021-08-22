package su.svn.cdc.utils;

// import ru.sberbank.bigdata.nrt.framework.core.util.sign.impl.SignatureCRC32Util;

import java.nio.ByteBuffer;

public class Crc32Util {
    public static ByteBuffer crc32(ByteBuffer byteBuffer) {
        byte[] bytes = new byte[byteBuffer.remaining()];
        byteBuffer.get(bytes).rewind();
        byte[] crc32 = null;// SignatureCRC32Util.sign(bytes);

        return ByteBuffer.wrap(crc32);
    }
}
