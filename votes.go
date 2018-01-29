// Everybody Votes Channel
// File generation
// Written in Golang by TheMrIron2 and Spotlight

package main

import "fmt"

type CountryVotingData struct {
    worldwide int32
    national int32
    national_results int32
    question_data *[]QuestionData // would have to be defined later on
    result_data *[]ResultData // likewise, as we figure it out
    country_code int32
    country_count int32
    language_code int32
    languages []int32 // I think this is multiple ints at least
    _ int32 // pretty sure num/number are unused
    _ int32
    nw string
    worldwide_q int32 // even if bool, needs to be written as 4 bytes
    national_q int32
    file_type // todo figure out xd
    write_questions int32
    write_results int32
} // thanks spot

fmt.Println("EVC File Generator") 
fmt.Println("By TheMrIron2 and Spotlight")
// you know honestly why is it even important to print this
// I digress

func time_convert(time int) int {
  return ((time - 946684800) / 60)
}

// use time.Now().Unix() instead of fetching epoch

func get_timestamp(mode, type, date) int {
    if mode == 0 {
        timestamp = time_convert(time.Now().Unix())
    else if mode == 1 or mode == 2 {
        timestamp = time_convert(time.mktime(date.timetuple()))
    }
        if mode == 2 {
            if production {
                if type == "n" {
                    timestamp += 10080 }
                else if type == "w" {
                    timestamp += 21600 }
            else {
                timestamp += 5 }
    return timestamp }}}

// ex. usage: get_timestamp(1, "n", get_date(q))

func days_ago() int {
    if national_results > 0 {
        return 7
    }
    else if national_results > 0 {
        return 14
    }
    else {
        return 0
    }
}

//year, month, day := fromDate.Date()
   
func get_poll_id() int {
    return poll_id
}
    
func pad(amnt) int {
    return "\0" * amnt // This code is garbage my god
}

func prepare() int {
    var country_count int32
    var countries int32
    var file_type int32
    var questions int32
    var poll_id int32
    var write_questions int32
    var write_results int32
    var results int32
    var position int32
    var national int32
    var worldwide int32
    
    /*
    you know as much as I dislike python and think Go is just plain better
    it is now also apparent that Go can also suck a cactus when it acts bad like this
    */
    
    fmt.Println("Preparing ...")
    if production { // todo: learn go, i have definitely just done something wrong
        client = Client(sentry_url)
        handler = SentryHandler(client)
        setup_logging(handler)
        logger = logging.getLogger(__name__)
    mysql_connect()
    }
    if len(sys.argv) == 1 {
        manual_run()
    }
    else if len(sys.argv) >= 2 {
        file_type = sys.argv[1]
    }
    else if file_type == "q" {
        automatic_questions()
    }
    else if file_type == "r" {
        automatic_results()
    }
    else if file_type == "v" {
        automatic_votes()
    }
    mysql_close()
    make_language_table()
}
    
/*
well the python docstring here is """Manually enter in what file type you want if no arguments are specified."""
however I don't know Go well enough to hack-fuck docstrings with actual code
sorry
just going to wait for snoot i suppose
this code is broken anyway lmao
*/
