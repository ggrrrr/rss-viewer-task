require "test_helper"

class SourcesControllerTest < ActionDispatch::IntegrationTest
  test "should get index" do
    get sources_index_url
    assert_response :success
  end

  test "should get delete" do
    get sources_delete_url
    assert_response :success
  end

  test "should get create" do
    get sources_create_url
    assert_response :success
  end
end
