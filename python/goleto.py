from cffi import FFI
from os import path

ffi = FFI()
ffi.cdef("""
    void RandomBoleto(char* out);
""")

if path.exists("./libgoleto.so"):
    lib = ffi.dlopen("./libgoleto.so")
elif path.exists("./goleto.dll"):
    lib = ffi.dlopen("./goleto.dll")

out = ffi.new("char[45]")

def random_boleto():
    lib.RandomBoleto(out)

    return ffi.string(out).decode('ascii')

