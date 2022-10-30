def isPalindrome(string):
    left, right = 0, len(string)-1
    while left < right:
        if string[left] != string[right]:
            return False
        left += 1
        right -= 1

    return True

if __name__ == "__main__":
    s = "abccba"
    print("Is palindrome?", isPalindrome(s))

# output: ('Is palindrome?', True)