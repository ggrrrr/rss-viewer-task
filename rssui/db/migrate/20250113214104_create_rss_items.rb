class CreateRssItems < ActiveRecord::Migration[8.0]
  def change
    create_table :rss_items do |t|
      t.text :uniq_hash
      t.text :title
      t.text :link
      t.text :source
      t.text :source_url
      t.text :description
      t.datetime :publish_date

      t.timestamps
    end
  end
end
