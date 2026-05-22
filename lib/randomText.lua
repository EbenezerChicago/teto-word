local dailyWord = arg[1]
local result = ""

if not dailyWord then
    result = "ERROR"
elseif type(arg[1]) == "table" then
    if os.date("%m") == "04" and os.date("%d") == "01" then
        result = "CELEBRANT!!!"
    else
        result = dailyWord[math.random(#dailyWord)]:upper()
    end
else
    result = arg[1]:upper()
end

return result