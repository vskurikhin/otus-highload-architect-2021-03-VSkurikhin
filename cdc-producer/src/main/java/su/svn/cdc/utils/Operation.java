package su.svn.cdc.utils;

/**
 * Enum to derive the operation that was performed in the DataBase.
 */
public enum Operation {

    READ("r"),
    CREATE("c"),
    UPDATE("u"),
    DELETE("d");

    private final String code;

    Operation(String code) {
        this.code = code;
    }

    public String code() {
        return this.code;
    }

    public static Operation forCode(String code) {
        Operation[] values = values();

        for (Operation op : values) {
            if (op.code().equalsIgnoreCase(code)) {
                return op;
            }
        }
        return null;
    }
}
