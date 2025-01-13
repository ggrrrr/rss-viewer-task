require 'uri'
require 'net/http'


class RssItem
  attr_accessor :title, :link, :source, :source_url, :description, :publish_date

  def initialize(title, link, source, source_url, description, publish_date)
    @title = title
    @link = link
    @source = source
    @source_url = source_url
    @description = description
    @publish_date = publish_date
  end

end

class ViewerController < ApplicationController
  def refresh
    RssJob.perform_later
  end

  def index
    urls = ["https://news.google.com/rss/search?hl=en-US&gl=US&q=samsung&um=1&ie=UTF-8&ceid=US:en"]

    array = Address.all
    array.each { |x| 
      urls.append(x.url)
    }
  
    uri = URI('http://localhost:8080/v1/parse')
    body = {urls:urls}
    request = Net::HTTP::Post.new(uri)
    headers = { 'Content-Type': 'application/json' }
    response = Net::HTTP.post(uri, body.to_json, headers)

    j = JSON.parse(response.read_body)

    @data = j['items'].inject([]) { |o,d| o << RssItem.new(
      d['title'],
      d['link'],
      d['source'],
      d['source_url'],
      d['description'],
      d['publish_date'].to_datetime
      )
    }
  end
end
