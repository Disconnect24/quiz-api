print("EVC File Generation but in Lua")
print("By TheMrIron2 for Disconnect24 (disconnect24.xyz)")

local worldwide = 0
local national = 0
local national_results = 0
local worldwide_results = 0
local question_data = {} -- should I use tables? No idea
local results = {}
local country_code = 49
local country_count = 0
local language_code = 1
local languages = {}
local num = 0 -- what the fuck?
local number = 0 -- what the fuck?
local nw = ""
local worldwide_q = false
local national_q = false
local file_type = 0 -- what the fuck?
local write_questions = false
local write_results = false

--[[
Note about modes
Mode 0 = Wii
Mode 1 = vWii, as far as I'm aware
Mode 2 = Likely debugging
--]]

local function time_convert(time)
  (time - 946684800) / 60
end

local function get_epoch()
  time() -- i think??????
end

function get_timestamp(mode, type, date)
    if mode == 0 then
        timestamp = time_convert(time())
    end
    else if mode == 1 or mode == 2 then
        timestamp = time_convert(time.mktime(date.timetuple()))
    end
        if mode == 2 then
            if production then
                if type == "n" then
                    timestamp += 10080
                    end
                  end
                else if type == "w" then
                    timestamp += 21600
                end
            else
            timestamp += 5
            end
end
end
end
