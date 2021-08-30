
math.randomseed(os.time())
math.random(); math.random(); math.random()

userNames = {'root', 'user', 'name', 'koch2804', 'lang8129', 'kihn1299', 'hahn1712', 'kris8365', 'nolan7356', 'auer2414'}

token = 'eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzAzNjU3NDEsImp0aSI6IjkzMzYwNjM2OTE3MzA0Mzc3OTAifQ.WzVFOa_sJ5_FtvIqpGSW4WBjCqVzBswwQUwgPakHESB6ouvHmqZpjgh1CiMTbmXyA5rIOQsc8nk6_YLl3tMjzA'

wrk.method = "POST"
wrk.body   = "foo=bar&baz=quux"
wrk.headers["Content-Type"] = "application/x-www-form-urlencoded"

request = function()
  rangeUserNames = table.getn(userNames) - 1
  username = userNames[math.random(rangeUserNames)]
  path = "/message"
  -- Return the request object with the current URL path
  wrk.body = '{"Message": "test' .. math.random(9999999) .. '", "ToUser": "' .. username .. '"}'
  return wrk.format(
          'POST',
          path,
          {
            ['Host'] = 'localhost',
            ["Content-Type"] = "application/x-www-form-urlencoded",
            ["Cookie"] = "acs_jwt=" .. token .. ";"
          })
end

response = function(status, headers, body)
  for key, value in pairs(headers) do
    if key == "Location" then
      io.write("Location header found!\n")
      io.write(key)
      io.write(":")
      io.write(value)
      io.write("\n")
      io.write("---\n")
      break
    end
  end
end
