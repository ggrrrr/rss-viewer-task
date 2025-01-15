class CreateRssSources < ActiveRecord::Migration[8.0]
  def change
    create_table :rss_sources do |t|
      t.string :url

      t.timestamps
    end
  end
end
