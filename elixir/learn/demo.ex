defmodule Genmetry do 
    def run do 
        IO.puts "Hello"
    end
    
    def add(a,b), do: a + b

    def add(a,b,c),do: a + b + c

    defp p_add(a,b) do 
        a + b 
    end 

    def default_argument(a,b \\ 1) do 
        a + b 
    end 

    import IO 
    def pipeline do 
        -5
        |> abs
        |> Integer.to_string 
        |> puts 
    end 

end 


# Genmetry.run()

# res = Genmetry.add 1,2
# IO.puts res

# Genmetry.pipeline()

# res = Genmetry.add(1,2,3)
# IO.puts res

# res = Genmetry.default_argument(1)
# IO.puts res 



defmodule Circle do 
    @moduledoc "Implement basic circle functions"

    @pi 3.14159 

    @spec area(number) :: number 
    def area(r), do: r * r * @pi 

    @spec circumference(number) :: number 
    def circumference(r),do: 2 * r * @pi 

end 

IO.puts Circle == :"Elixir.Circle"


defmodule OurMacro do 
    defmacro unless(expr, do: block) do 
        quote do 
            if !unquote(expr), do: unquote(block)
        end 
    end
end 


run_query = fn(query_def) -> 
    :timer.sleep(2000)
    "#{query_def} result "
end 


async_query = fn(query_def) -> 
    caller = self 
    spawn(fn -> 
        send(caller, {:query_result, run_query.(query_def)})
    end)
end 

Enum.each(1..5, &async_query.("query #{&1}"))

get_result = fn -> 
    receive do 
        {:query_result, result} -> result 
    end 
end 

results = Enum.map(1..5, fn(_) -> get_result.() end)

IO.puts "The result is: #{results}"

1..5
|> Enum.map(&async_query.("query #{&1}"))
|> Enum.map(fn(_) -> get_result.() end)
|> IO.puts 


defmodule DatabaseServer do 
    
    defp loop do 
        receive do 
            {:run_query, caller, query_def} -> 
                send(caller, {:query_result, run_query(query_def)})
        end 
        loop
    end 
    
    defp run_query(query_def) do 
        :timer.sleep(2000)
        "#{query_def} result"
    end 

end 
        