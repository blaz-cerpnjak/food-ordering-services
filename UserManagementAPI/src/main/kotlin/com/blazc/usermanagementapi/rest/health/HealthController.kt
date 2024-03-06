package com.blazc.usermanagementapi.rest.health

import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.CrossOrigin
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RestController
import java.net.InetAddress

@CrossOrigin
@RestController
class HealthController {

    @GetMapping("/health")
    fun getServiceInfo(): ResponseEntity<String> {
        var location = "??"

        try {
            val localHost = InetAddress.getLocalHost()
            location = "${localHost.hostName}:${localHost.hostAddress}"
        } catch (e: Exception) {
            e.printStackTrace()
            //log.info("Unable to get local address: ${e.message}")
        }

        return ResponseEntity.ok("$location->${this::class.java}")
    }

}