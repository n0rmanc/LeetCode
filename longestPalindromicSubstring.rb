# @param {String} s
# @return {String}
def longest_palindrome(s)
  results = {}
  max = 0
  s.length.times do |index|
    start_pop = false
    left_array = []
    find_palindrome = ''
    # puts "test string #{s[index..-1]}"
    s[index..-1].each_char do |char|
      if left_array.last == char
        start_pop = true
        find_palindrome += left_array.pop
        find_palindrome.prepend char
        # puts "AAAA #{char} #{left_array} #{start_pop} #{find_palindrome}"
      elsif left_array[-2] == char && !start_pop
        start_pop = true
        find_palindrome += left_array.pop
        find_palindrome += left_array.pop
        find_palindrome.prepend char
        # puts "BBBBB #{char} #{left_array} #{start_pop} #{find_palindrome}"
      else
        if start_pop
          start_pop = false
          left_array = []
          if find_palindrome.length > max
            results[find_palindrome.length] = find_palindrome
            max = find_palindrome.length
          end
          find_palindrome = ''
        end
        left_array.push char
        # puts "CCC #{char} #{left_array} #{start_pop}"
      end
      if find_palindrome.length > max
        results[find_palindrome.length] = find_palindrome
        max = find_palindrome.length
      end
    end
  end
  # puts "answer #{results[max]}"
  puts "answer #{results[max]}"
  results[max]
end

test = 'babad'
longest_palindrome(test)
test = 'abcdefedcba'
longest_palindrome(test)
test = 'abcdefedcbababad'
longest_palindrome(test)
test = 'bbbbbbbbbbbbbbbb'
longest_palindrome(test)


# 'abcdefedcba'
# a b c d e f e d c b a
