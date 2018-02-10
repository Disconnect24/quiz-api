// Everybody Votes Channel
// File generation
// Written in Golang by TheMrIron2 and Spotlight

package main

import "fmt"
import "database/sql"

type CountryVotingData struct {
    worldwide int32
    national int32
    national_results int32
    question_data *[]QuestionData // would have to be defined later on
    result_data *[]ResultData // will be figured out also
    country_code int32
    country_count int32
    language_code int32
    languages []int32 // what the fuck? [probably multiple ints]
    _ int32 // what the fuck?
    _ int32 // what the fuck?
    nw string // what the fuck?
    worldwide_q int32 // even if bool, needs to be written as 4 bytes
    national_q int32 
    file_type // what the fuck?
    write_questions int32
    write_results int32
}

fmt.Println("EVC File Generator") 
fmt.Println("By TheMrIron2 and Spotlight")
// you know honestly why is it even important to say this

func time_convert(time int) int {
  return ((time - 946684800) / 60) // what the fuck?
}

// use time.Now().Unix() instead of fetching epoch

func get_timestamp(mode, type, date) int {
    if mode == 0 {
        timestamp = time_convert(time.Now().Unix())
    else if mode == 1 or mode == 2 {
        timestamp = time_convert(time.mktime(date.timetuple())) // what the fuck?
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
    return poll_id // what the fuck?
} // why??
    
func pad(amnt) int {
    return "\0" * amnt // what the fuck?
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
    
    // does this even work?
    
    fmt.Println("Preparing ...")
    if production { // todo: learn go
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
    
// Manually enter in what file type you want if no arguments are specified.

func manual_run() int {
    question_count := len(question_data)
    fmt.Println("Loaded %s %s" % (question_count, "Question(s)"))
    file_type = raw_input('Enter File Type (q/r/v): ')
    if file_type == "q" {
        write_questions = true
    }
    else if file_type == "r" {
        write_results = true
    }
    else if file_type == "v" {
        if raw_input('Write Questions? (y/n): ') == "y" {
            write_questions = true
        }
        if raw_input('Write Results? (y/n): ') == "y" {
            write_results = true
        }
    }
    else {
        fmt.Println("Error: Invalid file type selected")
        os.Exit(1) // error code 1 indicates invalid file type for future reference
    }
    if file_type == "r" or (file_type == "v") { // (file_type == "v" and write results)?
        poll_id = int(raw_input('Enter Result Poll ID: ')) // Python-esque garbage
    }
}
    
// Automatically run the scripts. This will be what the crontab uses.


func automatic_questions() int {
    // defining vars that were shittily, globally defined in 1 line in Python in old code
    var write_questions int32
    var write_results int32
    var questions int32
    var nw int32
    
    write_questions = true
    nw = sys.argv[2]
    mysql_get_questions(1, nw) // todo: mysql in go
    questions = 1
}

func automatic_results() int {
    var write_results int32
    var results int32
    var national int32
    var worldwide int32
    var questions int32
    var national_results int32
    var worldwide_results int32
    var nw int32
    
    // sigh
    
    
    write_results = true
    nw = os.Args[2:]
    if nw == "n":
        days = 7
    else if nw == "w":
        days = 15
    results[get_poll_id()] = mysql_get_votes(days, nw, 1)
    delete results[None] // except keyerror pass??
    national = 0
    worldwide = 0
    questions = 0
}
    
func automatic_votes() int {
    var write_questions int32
    var write_results int32
    var questions int32
    var results int32
    var national int32
    var worldwide int32
    var questions int32
    
    write_questions = true
    write_results = true
    mysql_get_questions(1, "w")
    mysql_get_questions(3, "n")
    questions = national + worldwide
    question_count = len(question_data)
    printf("Loaded %s %s" % (question_count, "Questions"))
    for v in list(reversed(range(1, 7))) {
        results[get_poll_id()] = mysql_get_votes(7, "n", v)
        results[get_poll_id()] = mysql_get_votes(15, "w", 1)
    } // THIS DOESNT EVEN WORK because i don't know go
    delete results[None]
}
   
