defmodule Mix.Tasks.Example do 
    use Mix.Task

    @shortdoc "Simply runs the Hello.say/0 command"
    def run(_) do
        Example.say
    end
end
