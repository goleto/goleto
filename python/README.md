# Uso no python

Para usar esse projeto no python é necessário primeiro exportar um shared library, que pode ser importada usando o módulo cff.

Para isso, siga os seguintes passos:

1. Compilar a biblioteca

    ```sh
    # No windows
    go build -o goleto.dll -buildmode=c-shared main.go

    # No Linux ou Mac
    go build -o libgoleto.so -buildmode=c-shared main.go
    ```

2. Instalar o pacote `cffi`

    ```
    pip install cff
    ```

3. Usar a módulo `goleto` que está nesse diretório

    ```py
    from goleto import random_boleto

    print(random_boleto())
    ```