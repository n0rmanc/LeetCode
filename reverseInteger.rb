# @param {Integer} x
# @return {Integer}
def reverse(x)
  fixnum_max = 1<<31
  fixnum_min = 1<<31 * -1

  x_string = x.to_s
  is_negative = false
  if x < 0
    is_negative = true
    x_string[0] = ''
  end
  result = x_string.reverse.to_i
  result = result * -1 if is_negative
  puts "#{fixnum_max} #{fixnum_min} #{result} #{result > fixnum_max || result < fixnum_min}"
  return 0 if result > fixnum_max || result < fixnum_min
  result
end
reverse(0)
reverse(123)
reverse(-123)
reverse(2147483648)
reverse(-2147483412)
reverse(1534236469)
