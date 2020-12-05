def f
  raise
end

begin
  f
  puts 'Am I dead?'
rescue RuntimeError
  # do nothing
end
