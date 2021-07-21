box.cfg {
    listen = 3301;
    memtx_memory = 768 * 1024 * 1024; -- 128Mb
    memtx_min_tuple_size = 16;
    memtx_max_tuple_size = 768 * 1024 * 1024; -- 128Mb
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
        s:create_index('primary', {type = 'TREE', parts = {1, 'unsigned'}, if_not_exists = true})
    end

    if not box.space.user then
        t = box.schema.space.create('user')
        t:create_index('name', {type = 'TREE', unique = false, parts = {2, 'string'}, if_not_exists = true})
    end

end

fun = require('fun')

function user_index_name_search(a, b)
    result = {}
    rname = '^' .. a
    rsurname = '^' .. b
    items = box.space.user.index.name:select({a}, {iterator = 'GE'})
    for i = 1,#items do 
        if not string.match(items[i][2], rname) then break end
        if string.match(items[i][4], rsurname) then
            table.insert(result, items[i])
        end
    end
    return result
end

function user_index_name_search_f(a, b)
    return fun.iter(box.space.user.index.name:select({a}, {iterator = 'GE'}))
              :filter(function (tuple) return string.match(tuple[2], '^' .. a) end)
              :filter(function (tuple) return string.match(tuple[4], '^' .. b) end):totable()
end

bootstrap()
