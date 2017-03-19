# @param {String} s
# @return {String}
@left = -1
@right = -1
@max
def longest_palindrome(s)
  @left = -1
  @right = -1
  @max = -1
  return s if s.length < 2
  max = -1
  s.length.times do |index|
    check_palindrome(index, index, s, max)
    check_palindrome(index, index + 1, s, max)
  end
  s[@left, @max]
end

def check_palindrome(left, right, s, max)
  while left >= 0 && right <= s.length && s[left] == s[right]
    left -= 1
    right += 1
  end
  if @max < right - left - 1
    @left = left + 1
    @max =  right - left - 1
  end
end
