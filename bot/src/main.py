import os
import telebot
from telebot import types
from flask import Flask, request
import os 
import requests
from tokenData import token

Env = "Debug"
TOKEN = token
# TOKEN = os.getenv('TOKEN')
bot = telebot.TeleBot(TOKEN)

bot.set_my_commands([
    types.BotCommand('/start', "Start bot"),
])

@bot.message_handler(commands=['start'])
def send_welcome(commands):
    bot.send_message(message.chat.id, "Hello! Here you can create subscribes on your channel or subscribe on another channels!")

if __name__ == "__main__":
    if Env == "Debug":
        bot.remove_webhook()
        print("Bot started his working")
        bot.polling(none_stop=True)
