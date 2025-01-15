require "test_helper"

class RssSourcesControllerTest < ActionDispatch::IntegrationTest
  test "should get index" do
    get rss_sources_index_url
    assert_response :success
  end

  test "should get delete" do
    get rss_sources_delete_url
    assert_response :success
  end

  test "should get create" do
    get rss_sources_create_url
    assert_response :success
  end
end
