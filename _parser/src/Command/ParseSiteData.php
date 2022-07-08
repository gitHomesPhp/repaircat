<?php

declare(strict_types=1);

namespace App\Command;

use App\Entity\Url;
use App\Service\GisParser\Review\ReviewParser;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;

#[AsCommand(name: 'parse:site:data')]
class ParseSiteData extends Command
{
    private ReviewParser $reviewParser;

    public function __construct(
        private EntityManagerInterface $em,
    ) {
        $this->reviewParser = new ReviewParser($em);

        parent::__construct();
    }

    protected function execute(InputInterface $input, OutputInterface $output): int
    {
        $batchSize = 20;
        $i = 1;
        $q = $this->em->createQuery('SELECT u FROM ' . Url::class . ' u WHERE u.is_checked = false');

        foreach ($q->toIterable() as $url) {
            /**
             * @var $url Url
             */
            $this->reviewParser->handle($url);
            $i++;
            if (($i % $batchSize) === 0) {
                $this->em->flush();
                $this->em->clear();
            }
        }

        $this->em->flush();

        return Command::SUCCESS;
    }
}