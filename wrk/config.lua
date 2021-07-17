
math.randomseed(os.time())
math.random(); math.random(); math.random()

firstNames = {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
surNames = {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

request = function()
  rangeFirstNames = table.getn(firstNames) - 1
  rangeSurNames = table.getn(surNames) - 1
  path = "/users/search/" .. firstNames[math.random(rangeFirstNames)] .. "/" .. surNames[math.random(rangeSurNames)]
  -- Return the request object with the current URL path
  return wrk.format('GET', path, {['Host'] = 'localhost', ["Cookie"] = ""})
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
