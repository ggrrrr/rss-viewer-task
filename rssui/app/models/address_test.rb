require "test_helper"

class AddressTest < ActiveSupport::TestCase

    test "should not save address without url" do
        # assert true
        address = Address.new
        assert_not address.save, "Saved the address without a url"
      end

      test "should save address" do
        # assert true
        address = Address.new(url: "test_url")
        assert address.save, "Saved the address without a url"
      end
    
end

