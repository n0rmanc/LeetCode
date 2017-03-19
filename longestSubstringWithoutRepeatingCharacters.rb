# @param {String} s
# @return {Integer}
def length_of_longest_substring(s)
  max = 0
  patterns = {}
  s.length.times do |index|
    accepted_characters = []
    next_character_index = s[(index + 1)..-1].index(s[index])
    next_character_index = next_character_index.nil? ? s.length : next_character_index + index
    if !patterns[s[index..(next_character_index - 1)]].nil? ||
      s[index..(next_character_index - 1)].length <= max
      next
    end
    s[index..(next_character_index - 1)].each_char do |c|
      if accepted_characters.include? c
        break
      else
        accepted_characters.push c
      end
    end

    max = accepted_characters.count if accepted_characters.count > max
    patterns[a.join] = max
  end
  max
end
