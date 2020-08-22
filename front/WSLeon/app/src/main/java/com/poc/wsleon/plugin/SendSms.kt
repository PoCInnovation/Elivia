package com.poc.wsleon.plugin

import android.Manifest
import android.app.Activity
import android.content.Context
import android.content.pm.PackageManager
import android.telephony.SmsManager
import androidx.core.app.ActivityCompat

class SendSms(context:Context, activity: Activity) {
    private val mainContext: Context = context
    private val mainActivity: Activity = activity

    fun run(phoneNumber: String, message: String) {

        if (ActivityCompat.checkSelfPermission(mainContext, Manifest.permission.SEND_SMS) == PackageManager.PERMISSION_DENIED) {
            val requestSendSms: Int = 2

            ActivityCompat.requestPermissions(mainActivity, arrayOf(Manifest.permission.SEND_SMS), requestSendSms)
        } else{
            SmsManager.getDefault().sendTextMessage(phoneNumber, null, message, null, null)
        }
    }
}