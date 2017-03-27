
def convert(s, num_rows)
  sb = []
  num_rows.times { |i| sb[i] = '' }
  len = s.length
  i = 0
  while i < len
    idx = 0
    num_rows.times do |row|
      unless s[i].nil?
        sb[row] += s[i]
        i+=1
      end
    end
    (num_rows - 2).downto(1).each do |row|
      unless s[i].nil?
        sb[row] += s[i]
        i+=1
      end
    end
  end
  sb.join('')
end

convert('PAYPALISHIRING', 3)
convert('PAYPALISHIRING', 5)
convert('PAYPALISHIRING', 7)
convert('ABC', 2)
convert('ABCD', 2)
