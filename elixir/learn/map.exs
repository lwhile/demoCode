ExUnit.start

defmodule MapTest do 
    use ExUnit.Case 

    def sample do 
        %{foo: "bar", baz: "quz"}
    end 

    test "Map.get" do 
        assert Map.get(sample, :foo) == "bar"
        assert Map.get(sample, :bazzz) == nil 
    end 

    test "[]" do 
        assert sample[:foo] == "bar"
        assert sample[:baaaa] == nil 
    end 

    test "." do 
        assert sample.foo == "bar"
        assert_raise KeyError, fn ->
            sample.baaaa
        end 
    end 

    test "Map.fetch" do 
        {:ok, val} = Map.fetch(sample, :foo)
        assert val == "bar"
        :error = Map.fetch(sample, :asd)
    end 

    test "Update map using pattern matching syntax" do 
        # Only way to update existing key
        assert %{sample | foo: "bob"} == %{foo: "bob", baz: "quz"}
    end 
end 