# Electronic Systems DSL

## ESystemsDSL

```
component MainComponent {
    in port int32 
}

// component functions implemented in esystemsdsl
component Fibonacci {
    in port int32 n;
    out port int32 fib;
    
    var int32 a = 0;
    var int32 b = 1;
    var int32 i = 0;
    
    component Adder adder_inst;
    
    // sequential
    for i < n {
        // parallel
        a = b;
        b = a + b;
    }
    
    fib = b;
    
    // implicit return
}

component Counter {
    in signal bool enable;

    var int32 n = 0;
    
    for {
        n = n + 1;
    }
    
    return on enable falling;
}

// interaction with component interfaces, implemented in pure vhdl rtl
interface Adder {
    in port int32 a;
    in port int32 b;
    out port int32 c;
    
    // return condition
    return on clk rising;
    
    // behaviour defined in vhdl
}
```

## ESystemsIR

```

```

## VHDL

```vhdl
component Fibonacci {
    in port int32 n;
    out port int32 fib;
    
    var int32 a = 0;
    var int32 b = 1;
    var int32 i = 0;
    
    component Adder adder_inst;
    
    // sequential
    for i < n {
        // parallel
        a = b;
        b = a + b;
    }
    
    fib = b;
    
    // implicit return
}

entity Fibonacci is

end entity

architecture behaviour of Fibonacci is

begin

    process(clk)
    
    begin
        if rising_edge(clk):
            if (i < n):
                a := b;
                b := a + b;
            end if;
        end if;
    end process;

end architecture
```

