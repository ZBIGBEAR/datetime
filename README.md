# 封装时间相关的操作

# 1.func AddSecond(t time.Time, second int64) time.Time
在指定时间上增加秒

# 2.func AddMinute(t time.Time, minute int64) time.Time
在指定时间上增加分

# 3.func AddHour(t time.Time, hour int64) time.Time
在指定时间上增加时

# 4.func AddDay(t time.Time, day int64) time.Time
在指定时间上增加天

# 5.func AddMonth(t time.Time, month int64) time.Time
在指定时间上增加月

# 6.func AddYear(t time.Time, year int64) time.Time
在指定时间上增加年

# 7.func GetNowDate() string
获取当前日期，格式为：yyyy-mm-dd

# 8.func GetNowTime() string
获取当前时间，格式为：hh-mm-ss

# 9.func GetNowDateTime() string
获取当前日期时间，格式为：yyyy-mm-dd hh-mm-ss

# 10.func GetNowUnix() int64
获取当前时间戳

# 11.func TimeToStr(t time.Time, format DateTimeFormatType) string
将指定时间装换成指定格式的字符串

# 12.func StrToTime(str string, format DateTimeFormatType) (time.Time, error)
将字符串格式的世界转换成指定格式的时间

# 13.func BeginOfMinute(t time.Time) time.Time
获取当前日期时间的分钟的第一秒

# 14.func EndOfMinute(t time.Time) time.Time
获取当前日期时间的分钟的最后一秒

# 15.func BeginOfHour(t time.Time) time.Time
获取当前日期时间的小时的第一秒

# 16.func EndOfHour(t time.Time) time.Time
获取当前日期时间的小时的最后一秒

# 17.func BeginOfDay(t time.Time) time.Time
获取当前日期时间的天的第一秒

# 18.func EndOfDay(t time.Time) time.Time
获取当前日期时间的天的最后一秒

# 19.func BeginOfWeek(t time.Time) time.Time
获取当前日期时间的星期的第一秒

# 20.func EndOfWeek(t time.Time) time.Time
获取当前日期时间的星期的最后一秒

# 21.func BeginOfMonth(t time.Time) time.Time
获取当前日期时间的月的第一秒

# 22.func EndOfMonth(t time.Time) time.Time
获取当前日期时间的月的最后一秒

# 23.func BeginOfYear(t time.Time) time.Time
获取当前日期时间的年的第一秒

# 24.func EndOfYear(t time.Time) time.Time
获取当前日期时间的年的最后一秒