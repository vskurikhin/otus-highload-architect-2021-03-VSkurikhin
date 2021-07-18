box.cfg {
    listen = 3301;
    memtx_memory = 128 * 1024 * 1024; -- 128Mb
    memtx_min_tuple_size = 16;
    memtx_max_tuple_size = 128 * 1024 * 1024; -- 128Mb
    force_recovery = true;

     -- 1 – SYSERROR
     -- 2 – ERROR
     -- 3 – CRITICAL
     -- 4 – WARNING
     -- 5 – INFO
     -- 6 – VERBOSE
     -- 7 – DEBUG
     log_level = 6;
 }

local function bootstrap()

    if not box.space.mysqldaemon then
        s = box.schema.space.create('mysqldaemon')
        s:create_index('primary',
        {type = 'tree', parts = {1, 'unsigned'}, if_not_exists = true})
    end

    if not box.space.mysqldata then
        t = box.schema.space.create('user')
        t:create_index('primary',
        {type = 'tree', parts = {1, 'unsigned'}, if_not_exists = true})
    end

    if not box.space.mysqldata then
        t = box.schema.space.create('name_user_id')
        t:create_index('primary',
        {type = 'tree', parts = {1, 'string'}, if_not_exists = true})
    end

    if not box.space.mysqldata then
        t = box.schema.space.create('surname_user_id')
        t:create_index('primary',
        {type = 'tree', parts = {1, 'string'}, if_not_exists = true})
    end

end

bootstrap()
