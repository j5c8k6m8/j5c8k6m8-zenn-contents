module Kernel
  def raise
    puts 'override raise!!'
  end
end

begin
  raise
  puts 'Am I dead?'
rescue RuntimeError
end
