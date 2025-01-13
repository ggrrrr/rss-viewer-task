class RssJob < ApplicationJob
    queue_as :default

    def perform(*args)
      puts "Processing -> Fetching data from rss aggregator"

      # Fetch from API
      # Persist in ?
      # create models for RSS
    end

end