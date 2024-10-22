def main():
    with open("input/listing_0037_single_register_mov", "rb") as f:
        data = f.read(4)
        print(bin(int(data)))

if __name__ == "__main__":
    main()