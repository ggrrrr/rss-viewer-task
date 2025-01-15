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
    array = RssSource.all
    array.each { |x| 
      urls.append(x.url)
    }

    begin
      rssURI = URI(Rails.configuration.rss_url)
      jwt = URI(Rails.configuration.rss_jwt)
      body = {urls:urls}
      headers = { 'Content-Type': 'application/json','Authorization': "Bearer #{jwt}" }
      response = Net::HTTP.post(rssURI, body.to_json, headers)
      responseJson = JSON.parse(response.read_body)
      responseJson['items'].each { |d|
        save_item(d)
      }
    rescue => e
      logger.error e.message
      logger.error e.backtrace.join("\n")
    end
  end
end