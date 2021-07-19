
math.randomseed(os.time())
math.random(); math.random(); math.random()

firstNames = {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
surNames = {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

request = function()
  rangeFirstNames = table.getn(firstNames) - 1
  rangeSurNames = table.getn(surNames) - 1
  path = "/users/search/" .. firstNames[math.random(rangeFirstNames)] .. "/" .. surNames[math.random(rangeSurNames)]
  -- Return the request object with the current URL path
  return wrk.format('GET', path, {['Host'] = 'localhost', ["Cookie"] = "_ga=GA1.1.100769609.1603607547; acs_jwt=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjY3NjQ1NjYsImp0aSI6IjkzMjgzMzAyMjc0MzIyMjI5NjMifQ.tRIr7m-eBlY-Y0-oVsouXwHEJS4ABhruaB5G--Zqk6XMqVK70Qxf_0V2t0VvQsEq_JJzuRQ4dChD6R0dP3mekA"})
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
