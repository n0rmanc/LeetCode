# @param {String} s
# @param {String} p
# @return {Boolean}
def is_match(s, p)
  dp = []
  dp[0] = [true]
  another_array = []
  for i in 0..(p.length - 1)
    dp[0][i] = true if p[i] == '*'
  end
  for i in 0..(s.length - 1)
    for j in 0..(p.length - 1)
      dp[i + 1] ||= [false]
      dp[i + 1][j + 1] = dp[i][j] if s[i] == p[j]
      dp[i + 1][j + 1] = dp[i][j] if p[j] == '.'
      next unless p[j] == '*'
      if p[j - 1] != s[i] && p[j - 1] != '.'
        dp[i + 1][j + 1] = dp[i + 1][j - 1]
      else
        dp[i + 1][j + 1] = (dp[i - 1][j] || dp[i][j + 1] || dp[i + 1][j - 2])
      end
    end
  end
  dp.each do |ele|
    puts ele.to_s
  end
  # puts dp.to_s
  puts dp[s.length - 1][p.length - 1]
end

# is_match('aa', 'aa')
is_match('aaaaac', 'a*c')
# is_match('aaaaac', 'a*')
# is_match('a', '.')
# is_match('a', 'a*')
# is_match('aa', '.*')
# is_match('aaa', 'a*a')
# is_match("aa","aa")
# is_match("aaa","aa")
# is_match("aa", "a*")
# is_match("aa", ".*")
# is_match("ab", ".*")
# is_match('aab', 'c*a*b')
