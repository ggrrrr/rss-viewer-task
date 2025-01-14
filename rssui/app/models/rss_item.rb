class RssItem < ApplicationRecord
  scope :title_like, lambda {|title| {:conditions => ["title LIKE ?", "#{title}%"]}}
  scope :source_like, lambda {|source| {:conditions => ["source LIKE ?", "#{source}%"]}}
  # scope :desc, order(name: :desc)

  # default_scope { order(publish_date: :desc) }
    # default_scope { order(publish_date: :desc) }
  scope :desc, -> { order("publish_date desc") }
end
