
math.randomseed(os.time())
math.random(); math.random(); math.random()

token = 'eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mjk5NjA3NTUsImp0aSI6IjkzMzUxOTM5OTE4OTk2ODc2ODQifQ.6DARh7sBTnLxC9_-Tj4-5zTxupwqwrHGdMGUysICTNq35T6xpj-g7VmpvR0zsyt2aRIN2zj4q-zPkqXkn0V4Nw'

wrk.method = "POST"
wrk.body   = "foo=bar&baz=quux"
wrk.headers["Content-Type"] = "application/x-www-form-urlencoded"

request = function()
  path = "/message"
  -- Return the request object with the current URL path
  wrk.body = '{"Message": "test' .. math.random(9999999) .. '", "ToUser": "root"}'
  return wrk.format('POST', path, {['Host'] = 'localhost', ["Content-Type"] = "application/x-www-form-urlencoded", ["Cookie"] = "acs_jwt=" .. token .. ";"})
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
