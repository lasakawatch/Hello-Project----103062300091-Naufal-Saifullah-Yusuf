import keyboard
import time

# Setting the delay between clicks
delay = 0.1  # In seconds

# Press space key continuously
while True:
    keyboard.press("space")
    time.sleep(delay)
    keyboard.release("space")