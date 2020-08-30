package com.poc.wsleon.plugin.src

import android.Manifest
import android.app.Activity
import android.content.ContentResolver
import android.content.Context
import android.content.pm.PackageManager
import android.os.Build
import android.provider.ContactsContract
import android.telephony.SmsManager
import android.util.Log
import androidx.core.app.ActivityCompat
import androidx.core.app.ActivityCompat.requestPermissions
import com.poc.wsleon.ui.LeonView
import java.util.*
import kotlin.collections.HashMap

class SendSms(context: Context, activity: Activity, view: LeonView) {
    private val mainContext: Context = context
    private val mainActivity: Activity = activity
    private val chat = view

    fun run(contactName: String, message: String) {
        val phoneNumber: String = getPhoneNumberFromName(contactName)

        if (ActivityCompat.checkSelfPermission(mainContext, Manifest.permission.SEND_SMS) == PackageManager.PERMISSION_DENIED) {
            val requestSendSms: Int = 2

            requestPermissions(mainActivity, arrayOf(Manifest.permission.SEND_SMS), requestSendSms)
        } else {
            if (phoneNumber == "") {
                chat.addOliviaBubble("Sorry, but I could not find this contact.")
                return
            }
            SmsManager.getDefault().sendTextMessage(phoneNumber, null, message, null, null)
        }
    }

    private fun getPhoneNumberFromName(contactName: String): String {
        val contactList: LinkedList<Map<String, String>> = loadContactList()
        var phoneNumber = ""
        var tmp: HashMap<String, String> = HashMap<String, String>()

        for (i in 0 until contactList.size) {
            for (key in contactList[i].keys) {
                if (key == contactName) {
                    tmp = contactList[i] as HashMap<String, String>
                    phoneNumber = tmp[key].toString()
                    tmp.clear()
                }
            }
        }
        return phoneNumber
    }

    private fun loadContactList(): LinkedList<Map<String, String>> {
        var builder = LinkedList<Map<String, String>>()

        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.M && ActivityCompat.checkSelfPermission(mainContext, Manifest.permission.READ_CONTACTS) != PackageManager.PERMISSION_GRANTED) {
            requestPermissions(mainActivity, arrayOf(Manifest.permission.READ_CONTACTS), 1)
        } else {
            builder = getContacts()
        }
        return (builder)
    }

    private fun getContacts(): LinkedList<Map<String, String>> {
        val contactList: LinkedList<Map<String, String>> =  LinkedList<Map<String, String>>()
        val resolver: ContentResolver = mainContext.contentResolver;
        val cursor = resolver.query(ContactsContract.Contacts.CONTENT_URI, null, null, null, null)

        if (cursor != null) {
            if (cursor.count > 0) {
                while (cursor.moveToNext()) {
                    val id = cursor.getString(cursor.getColumnIndex(ContactsContract.Contacts._ID))
                    val name = cursor.getString(cursor.getColumnIndex(ContactsContract.Contacts.DISPLAY_NAME))
                    val phoneNumber = (cursor.getString(
                        cursor.getColumnIndex(ContactsContract.Contacts.HAS_PHONE_NUMBER))).toInt()

                    if (phoneNumber > 0) {
                        val cursorPhone = mainContext.contentResolver.query(
                            ContactsContract.CommonDataKinds.Phone.CONTENT_URI,
                            null, ContactsContract.CommonDataKinds.Phone.CONTACT_ID + "=?", arrayOf(id), null)

                        if (cursorPhone != null) {
                            if(cursorPhone.count > 0) {
                                while (cursorPhone.moveToNext()) {
                                    val phoneNumValue = cursorPhone.getString(
                                        cursorPhone.getColumnIndex(ContactsContract.CommonDataKinds.Phone.NUMBER))
                                    val map = HashMap<String, String>()
                                    map[name] = phoneNumValue
                                    contactList.addLast(map)
                                }
                            }
                        }
                        cursorPhone?.close()
                    }
                }
            } else {
                Log.d("Plugin", "SendSms: No contacts available.")
            }
        }
        cursor?.close()
        return contactList
    }
}