class RssJob < ApplicationJob
    queue_as :default
    def save_item(json_item)
      d = json_item
      item = RssItem.new
      item.title = d['title']
      item.link = d['link']
      item.source = d['source']
      item.source_url = d['source_url']
      item.description = d['description']
      item.publish_date = d['publish_date'].to_datetime

      if RssItem.exists?(title: d['title'])
        item = nil
      else
        item.save
      end
    end

    def perform(*args)
      urls = []
      array = Address.all
      array.each { |x| 
        urls.append(x.url)
      }
    
      rssURI = URI(Rails.configuration.rss_url)
      body = {urls:urls}
      headers = { 'Content-Type': 'application/json' }
      response = Net::HTTP.post(rssURI, body.to_json, headers)
      responseJson = JSON.parse(response.read_body)
      responseJson['items'].each { |d|
        save_item(d)
      }
    end

end