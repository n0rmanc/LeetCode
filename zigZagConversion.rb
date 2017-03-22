require 'pry-byebug'
# @param {String} s
# @param {Integer} num_rows
# @return {String}
def convert(s, num_rows)
  return s if s.nil? || s.empty?
  puts s
  results = []
  column_count = 0
  s.length.times do |index|
    column_position = index / num_rows
    row_position = index % (num_rows + 1)

    # row_position = num_rows / 2 if (index + 1) % (num_rows + 1) == 0


    if (index + 1) % (num_rows + 1) == 0
      num_rows.times do |row|
        results[row] << if row == num_rows / 2
                          puts "#{column_position} #{row} #{s[index]}"
                          s[index]
                        else
                          puts "#{column_position} #{row} ' '"
                          ' '
                        end
      end
    else
      puts "#{column_position} #{row_position} #{s[index]}"
      results[row_position] ||= ''
      results[row_position] << s[index]
    end
  end
  binding.pry
  puts results.to_s
  puts result2.strip
end

convert('PAYPALISHIRING', 3)
