package com.poc.wsleon.plugin.core

import android.app.Activity
import android.content.Context
import android.util.Log
import com.poc.wsleon.plugin.src.OpenApp
import com.poc.wsleon.plugin.src.SendSms
import com.poc.wsleon.ui.LeonView
import org.json.JSONObject

class PluginManager(context: Context, activity: Activity, view: LeonView) {
    private val sendSms = SendSms(context, activity, view)
    private val openApp = OpenApp(context)
    private val mainActivity = activity
    private val chat = view

    fun run(pluginName: String, data: JSONObject) {
        if (pluginName == "sms") {
            smsPluginManager(data)
        } else if (pluginName == "openapp") {
            openAppPluginManager(data)
        }
    }

    private fun smsPluginManager(data: JSONObject) {
        val contactName = data["contact"] as String
        val message = data["message"] as String

        sendSms.run(contactName, message)
    }

    private fun openAppPluginManager(data: JSONObject) {
        val appName = data["appname"] as String

        openApp.run(appName)
    }
}