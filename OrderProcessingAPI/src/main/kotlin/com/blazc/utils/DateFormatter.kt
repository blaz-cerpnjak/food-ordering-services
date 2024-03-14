package com.blazc.utils

import java.time.LocalDateTime

class DateFormatter {
    companion object {

        fun toLocalDateTime(date: String): LocalDateTime {
            val localDateTime = LocalDateTime.parse(date, java.time.format.DateTimeFormatter.ISO_DATE_TIME)
            return localDateTime
        }

        fun toString(date: LocalDateTime): String {
            val string = date.format(java.time.format.DateTimeFormatter.ISO_DATE_TIME)
            return string
        }

    }
}