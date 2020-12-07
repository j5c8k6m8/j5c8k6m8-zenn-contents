def f
  yield
  puts 'Am I dead?'
end

f do
  break
end
