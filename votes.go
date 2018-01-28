// Everybody Votes Channel
// File generation
// Written in Golang by TheMrIron2 and Spotlight

package main

type CountryVotingData struct {
    worldwide int32
    national int32
    national_results int32
    question_data *[]QuestionData // would have to befined later on
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

func days_ago() int { // what a useful function
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
