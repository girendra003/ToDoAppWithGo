package cronjobs

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func Ok() {
	fmt.Printf("Starting cron jobs...\n")
	c := cron.New()

	value := 0

	// Use a closure to capture and update `value`
	// c.AddFunc("* * * * *", Task1(&value))
	c.AddFunc("@every 1s", Task1(&value))
// CRON_EXPRESSION,your function 
	c.Start()

	// select {} // Keep the program running///
}

// 2. Predefined Schedules
// "@every 1s" — every 1 second
// "@hourly" — every hour
// "@daily" — every day
// "@weekly" — every week
// "@monthly" — every month
// "@yearly" — every year

// 	3. Custom Interval
// 	You can change "1s" to "5m", "2h", etc., for different intervals:

// 	"@every 5m" // every 5 minutes
// 	"@every 2h" // every 2 hours



// minute hour day month day-of-week
// 	 *      *    *   *    *

// 	4. Specific Times
// 	You can also specify exact times using the cron format:
// 	"0 0 * * *" // every day at midnight
// 	"0 12 * * *" // every day at noon
// 	"0 0 * * 0" // every Sunday at midnight
// 	"0 0 1 * *" // first day of every month at midnight
// 	"0 0 1 1 *" // first day of January at midnight (New Year's Day)
// 	"0 0 * * 1-5" // every weekday at midnight
// 	"0 0 * * 6,7" // every weekend at midnight (Saturday and Sunday)
// 	"0 0 1-7 * *" // first seven days of every month at midnight
// 	"0 0 1-7 * *" // first seven days of every month at midnight
// 	"0 0 * * 1-5" // every weekday at midnight
// 	"0 0 * * 6,7" // every weekend at midnight (Saturday and Sunday)
// 	"0 0 1-7 * *" // first seven days of every month at midnight
// 	"0 0 1-7 * *" // first seven days of every month at midnight
// 	"0 0 * * 1-5" // every weekday at midnight
// 	"0 0 * * 6,7" // every weekend at midnight (Saturday and Sunday)
// 	"0 0 1-7 * *" // first seven days of every month at midnight
// 	"0 0 1-7 * *" // first seven days of every month at midnight
// 	"0 0 * * 1-5" // every weekday at midnight
// 	"0 0 * * 6,7" // every weekend at midnight (Saturday and Sunday)
// 	"0 0 1-7 * *" // first seven days of every month at midnight
// 	"0 0 1-7 * *" // first seven days of every month at midnight
// 	"0 0 * * 1-5" // every weekday at midnight
// 	"0 0 * * 6,7" // every weekend at midnight (Saturday and Sunday)
// 	"0 0 1-7 * *" // first seven days of every month at midnight
// 	"0 0 1-7 * *" // first seven days of every month at midnight
// 	"0 0 * * 1-5" // every weekday at midnight
// 	"0 0 * * 6,7" // every weekend at midnight (Saturday and Sunday)
// 	"0 0 1-7 * *" // first seven days of every month at midnight
// 	"0 0 1-7 * *" // first seven days of every month at midnight
// 	"0 0 * * 1-5" // every weekday at midnight
// 	"0 0 * * 6,7" // every weekend at midnight (Saturday and Sunday)