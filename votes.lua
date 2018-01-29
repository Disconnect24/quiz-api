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

local function get_poll_id()
  return poll_id
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

local function days_ago()
  if national_results > 0 then
    return 7
  elseif worldwide_results > 0 then
    return 14
  else
    return 0
end
end
end

--[[
The original code has a bunch of functions
that return the month and day etc.
I'm thinking about making one function that
calculates the difference between days/months etc.
so we can do this in an alternative manner.

eg.
local timeDiff = function(t2,t1)
	local d1,d2,carry,diff = os.date('*t',t1),os.date('*t',t2),false,{}
	local colMax = {60,60,24,os.date('*t',os.time{year=d1.year,month=d1.month+1,day=0}).day,12}
	d2.hour = d2.hour - (d2.isdst and 1 or 0) + (d1.isdst and 1 or 0) -- handle dst
	for i,v in ipairs({'sec','min','hour','day','month','year'}) do 
		diff[v] = d2[v] - d1[v] + (carry and -1 or 0)
		carry = diff[v] < 0
		if carry then diff[v] = diff[v] + colMax[i] end
	end
	return diff
end
--]]

local function prepare()
    local country_count, countries, file_type, questions, poll_id, write_questions, write_results, results, position, national, worldwide -- what the fuck? i hate unclear variable defining
    print("Preparing ...)
    if production == true then
        client = Client(sentry_url)
        handler = SentryHandler(client)
        setup_logging(handler)
        logger = logging.getLogger(__name__)
    mysql_connect() -- Honestly I'm not even bothered to change this python connection code right now
    end
    if len(sys.argv) == 1 then 
      manual_run()
    elif len(sys.argv) >= 2 then
      file_type = sys.argv[1]
      end
      if file_type == "q" then
        automatic_questions()
        end
          elif file_type == "r" then
            automatic_results()
            end
          elif file_type == "v" then
            automatic_votes()
            end
            mysql_close()
            make_language_table()
end

-- Manually enter in what file type you want if no arguments are specified.

function manual_run()
    question_count = len(question_data)
    print "Loaded %s %s" % (question_count, if question_count == 1 then "Question" else "Questions") -- hnng why are you doing this in brackets here
    file_type = raw_input('Enter File Type (q/r/v): ')
    if file_type == "q" then
        write_questions = true
    end
    elseif file_type == "r" then
        write_results = true
    elseif file_type == "v" then
        if raw_input('Write Questions? (y/n): ') == "y": write_questions = true
        if raw_input('Write Results? (y/n): ') == "y": write_results = true
    end
    end
    else
        error("Invalid file type selected")
        end
    if file_type == "r" or (file_type == "v" and write_results) then
    poll_id = int(raw_input('Enter Result Poll ID: '))

-- """Automatically run the scripts. This will be what the crontab uses."""

--[[
reading the python code and trying to not only understand it but translate to Lua is painstaking
i'm gonna leave it here for a while
--]]
