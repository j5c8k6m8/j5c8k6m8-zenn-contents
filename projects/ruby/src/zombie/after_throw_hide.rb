def raise
  puts 'hide raise!!'
end

begin
  raise
  puts 'Am I dead?'
rescue RuntimeError
end
