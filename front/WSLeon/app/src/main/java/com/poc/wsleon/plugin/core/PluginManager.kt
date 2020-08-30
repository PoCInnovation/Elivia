package com.poc.wsleon.plugin.core

import android.app.Activity
import android.content.Context
import com.poc.wsleon.plugin.src.OpenApp
import com.poc.wsleon.plugin.src.SendSms
import com.poc.wsleon.ui.LeonView

class PluginManager(context: Context, activity: Activity, view: LeonView) {
    val sendSms = SendSms(context, activity, view)
    val openApp = OpenApp(context)
}