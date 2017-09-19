package db

const DAY_QUERY string = `SELECT avg(usd) as usd, avg(eur) as eur, avg(rur) as rur, type, moment
							FROM prices
							WHERE date(moment) >= date(current_date, '-1 day') AND
      						date(moment) <= current_date
							GROUP BY strftime("%H", moment), type;`
const MONTH_QUERY string = `SELECT avg(usd) as usd, avg(eur) as eur, avg(rur) as rur, type, moment
							FROM prices
							WHERE date(moment) >= date(current_date, '-30 day') AND
      						date(moment) <= current_date
							GROUP BY strftime("%D", moment), type;`
const WEEK_QUERY string = `SELECT avg(usd) as usd, avg(eur) as eur, avg(rur) as rur, type, moment
							FROM prices
							WHERE date(moment) >= date(current_date, '-7 day') AND
      						date(moment) <= current_date
							GROUP BY strftime("%D", moment), type;`

